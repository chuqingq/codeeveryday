-module(test_monitor).

-export([test/0]).

test() ->
    P = spawn(fun() -> receive Req -> io:format("Req: ~p~n", [Req]) end end),
    Ref = erlang:monitor(process, P),
    io:format("Ref: ~p~n", [Ref]),
    P ! ok,
    receive Res -> io:format("Res: ~p~n", [Res]) end,
    ok.
   