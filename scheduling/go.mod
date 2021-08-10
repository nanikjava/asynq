module github.com/hibiken/asynq/scheduling

go 1.16

replace scheduling/tasks => ./

replace github.com/hibiken/asynq => /home/nanik/Downloads/temp/packages/src/github.com/hibiken/asynq

require (
	github.com/hibiken/asynq v0.18.3
	scheduling/tasks v0.0.0-00010101000000-000000000000
)
