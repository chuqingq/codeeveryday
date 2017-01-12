-module(test_link).

-export([test/0]).

test() ->
    process_flag(trap_exit, true),
    P = spawn_link(fun() -> receive A -> io:format("A:~p~n", [A]) end end),
    P ! abc,
    receive Info -> io:format("Info: ~p~n", [Info]) end,
    ok.

