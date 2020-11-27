package runner

import (
	"context"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/vrischmann/envconfig"
	"go.uber.org/zap"
	"sigs.k8s.io/yaml"
)

// Manager implements template method pattern to orchestrate and execute runner functions in a proper order.
type Manager struct {
	runner         Runner
	cfg            Config
	log            *zap.Logger
	statusReporter StatusReporter
}

// NewManager returns new Manager instance.
func NewManager(runner Runner, statusReporter StatusReporter) (*Manager, error) {
	var cfg Config
	err := envconfig.InitWithPrefix(&cfg, "RUNNER")
	if err != nil {
		return nil, errors.Wrap(err, "while loading configuration")
	}

	log, err := getLogger(cfg.LoggerDevMode)
	if err != nil {
		return nil, errors.Wrap(err, "while creating zap logger")
	}

	log = log.Named("runner").Named(runner.Name())
	loggerInto(log, runner)

	return &Manager{
		runner:         runner,
		cfg:            cfg,
		log:            log,
		statusReporter: statusReporter,
	}, nil
}

// Execute underlying runner function in a proper order.
func (r *Manager) Execute(stop <-chan struct{}) error {
	runnerInputData, err := r.readRunnerInput()
	if err != nil {
		return errors.Wrap(err, "while reading runner input")
	}

	ctx, cancel := r.cancelableContext(stop, runnerInputData.Context.Timeout)
	defer cancel()

	log := r.log.With(zap.String("runner", r.runner.Name()), zap.Bool("dryRun", runnerInputData.Context.DryRun))
	log.Debug("Starting runner")
	sout, err := r.runner.Start(ctx, StartInput{
		ExecCtx: runnerInputData.Context,
		Args:    runnerInputData.Args,
	})
	if err != nil {
		return errors.Wrap(err, "while starting action")
	}
	log.Debug("Runner started", zap.Any("status", sout.Status))

	if err = r.statusReporter.Report(ctx, runnerInputData.Context, sout.Status); err != nil {
		return errors.Wrap(err, "while setting status")
	}

	log.Debug("Waiting for runner completion")
	wout, err := r.runner.WaitForCompletion(ctx, WaitForCompletionInput{ExecCtx: runnerInputData.Context})
	if err != nil {
		log.Error("while waiting for runner completion", zap.Error(err))
		return errors.Wrap(err, "while waiting for completion")
	}
	log.Debug("Runner job completed",
		zap.Bool("success", wout.Succeeded),
		zap.String("message", wout.Message),
	)

	return wout.ErrorOrNil()
}

func (r *Manager) readRunnerInput() (InputData, error) {
	rawInput, err := ioutil.ReadFile(r.cfg.InputPath)
	if err != nil {
		return InputData{}, errors.Wrap(err, "while reading input data from disk")
	}

	var input InputData
	if err := yaml.Unmarshal(rawInput, &input); err != nil {
		return InputData{}, errors.Wrap(err, "while unmarshaling input data")
	}

	return input, nil
}

// cancelableContext returns context that is canceled when stop signal is received or configured timeout elapsed.
func (r *Manager) cancelableContext(stop <-chan struct{}, timeout Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout.Duration())
	}

	go func() {
		select {
		case <-ctx.Done():
		case <-stop:
			cancel()
		}
	}()

	return ctx, cancel
}

// LoggerInjector is used by the Manager to inject logger to Runner.
type LoggerInjector interface {
	InjectLogger(*zap.Logger)
}

// loggerInto sets logger on `runner` if requested.
func loggerInto(log *zap.Logger, runner interface{}) {
	if s, ok := runner.(LoggerInjector); ok {
		s.InjectLogger(log)
	}
}

func getLogger(loggerDevMode bool) (*zap.Logger, error) {
	if loggerDevMode {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}