-module(test).

-compile(export_all).

test() ->
	test(65536).

test(0) -> ok;
test(N) ->
	{ok, S} = gen_tcp:connect("127.0.0.1", 8888, [binary, {active, false}]),
	ok = gen_tcp:close(S),
	io:format("~p~n", [N]),
	test(N-1).

