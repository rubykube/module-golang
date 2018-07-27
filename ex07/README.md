# Workers

Implement a Goroutines pool manager that controls maximum number of running
goroutines in pool.

Scheduler will read input data from stdin in interactive mode, scheduler must
convert each line of input to float, spawn worker (goroutine), and
sleep specified number of seconds in spawned goroutine.

The maximum number of workers must be limited by variable that passed to
function.

Let's say we run scheduler with max_workers=10, it means that scheduler can
spawn maximum 10 goroutines, but it must not spawn them just right after
start, only if there is a work to be done.

If number of received tasks is less than 10 then number of running
goroutines must be less than 10 too.

If user sent only 5 tasks, then number of running workers must be 5 only.

If user sent 12 tasks, then number of running workers must be 10 (because
max_workers=10) and then existing workers must be re-used for finishing last 2
tasks, and free 8 workers must be stopped because there are no tasks to be
done.

## Output

* Before spawning new worker scheduler must print message: `worker:$worker_id spawning`
* After starting executing new task worker must print message: `worker:$worker_id sleep:$sleep_seconds`
* Before stopping worker goroutine scheduler must print message: `worker:$worker_id stopping`

    Where $worker_id is number of worker starting from 1.
    Where $sleep_seconds is given number of seconds to sleep.

## Level: advanced

Provide HTTP API /stats and expose number of running goroutines at the time.

## Level: graceful

Scheduler must support graceful shutdown, it means that if daemon received
SIGTERM/SIGINT linux process signal it should stop accepting new tasks from
stdin (just close stdin file descriptor), wait for running tasks and stop
the process.

## Level: serious

The maximum number of workers will be specified in file and scheduler must
support reloading configuration in real-time, it means that if daemon received
SIGHUP linux process signal it should re-load configuration file and adjust
number of maximum running goroutines.

# Hints & Links


* Buffered Channels
    - [Unbuffered and buffered channels](https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/unbuffered-and-buffered-channels.html)
    - [Introduction To Golang: Buffered Channels](https://www.openmymind.net/Introduction-To-Go-Buffered-Channels/)
    - [Buffered Channels and Worker Pools](https://golangbot.com/buffered-channels-worker-pools/)
* Non-blocking channel operations
    - [Go by Exampel: Select](https://gobyexample.com/select)
    - [Go by Example: Non-Blocking Channel Operations](https://gobyexample.com/non-blocking-channel-operations)
* Wait Group
    - [GoDoc to sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup)
    - [Use sync.WaitGroup in Golang](https://nanxiao.me/en/use-sync-waitgroup-in-golang/)
