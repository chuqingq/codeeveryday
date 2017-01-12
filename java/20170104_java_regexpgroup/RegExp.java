package cqq_test.cqq_test;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class RegExp {
	public static void main(String[] args) {
		String arg = "account:password@192.168.1.1:1234/path/to/node";
		Pattern p = Pattern.compile("(\\w+):(\\w+)@([^/]+)(/.*)");
		Matcher m = p.matcher(arg);

		if (!m.find() || m.groupCount() != 4) {
			System.out.println("arg is invalid");
			return;
		}

		System.out.println("group(0): " + m.group(0));
		System.out.println("groupCount: " + m.groupCount());
		System.out.println("account: " + m.group(1));
		System.out.println("password: " + m.group(2));
		System.out.println("host:port: " + m.group(3));
		System.out.println("path: " + m.group(4));
	}

	// 输出的内容是
	// group(0): account:password@192.168.1.1:1234/path/to/node
	// groupCount: 4
	// account: account
	// password: password
	// host:port: 192.168.1.1:1234
	// path: /path/to/node
}
