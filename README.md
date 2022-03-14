# task-gamezop

There are 3 different go projects inside this folder.
1. server a
2. debouncer
3. server b

Requirements: golang, redis

# How to run.

After setting up redis(assuming it's running on its default port :6379 with no password and default db)

open 3 different terminal in task folder

In the first terminal run,
```
make run_server_a
```

In the second terminal run,
```
make run_debouncer
```

And in the third terminal run,
```
make run_server_b
```

Once all of these services are up, it's done. 

-server a runs on :8000

-server b runs on :9000

-You can change all these configs from the .env present in the individual project folder.
