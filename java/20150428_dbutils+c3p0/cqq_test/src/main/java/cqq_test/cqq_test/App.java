package cqq_test.cqq_test;

import java.sql.SQLException;
import java.util.List;

import org.apache.commons.dbutils.QueryRunner;
import org.apache.commons.dbutils.handlers.BeanListHandler;

/**
 * Hello world!
 *
 */
public class App {
	public static void main(String[] args) throws Exception {
		DBHelper.initDB();
		String sql = "SELECT * FROM config";
		QueryRunner r = DBHelper.createQueryRunner();
		try {
			List<Config> configs = r.query(sql, new BeanListHandler<Config>(
					Config.class));
			System.out.print(configs);
		} catch (SQLException ex) {
			ex.printStackTrace();
		}
	}
}
