package cqq_test.cqq_test;

import java.beans.PropertyVetoException;
import java.sql.Connection;
import java.sql.SQLException;

import org.apache.commons.dbutils.QueryRunner;
import org.apache.log4j.Logger;

import com.mchange.v2.c3p0.ComboPooledDataSource;

public class DBHelper {
	private static Logger LOGGER = Logger.getLogger(DBHelper.class);

	private static ComboPooledDataSource dataSource;

	public static void initDB() throws ClassNotFoundException, SQLException {
		LOGGER.info("initDB");

		// dbUrl
		String dbUrl = "jdbc:mysql://127.0.0.1:3306/beegather?createDatabaseIfNotExist=true&characterEncoding=utf8&autoReconnect=true&failOverReadOnly=false";

		// dbUser
		String dbUser = "root";

		// dbPassword
		String dbPassword = "Xianchang88";

		// c3p0
		try {
			dataSource = new ComboPooledDataSource();
			dataSource.setUser(dbUser);
			dataSource.setPassword(dbPassword);
			dataSource.setJdbcUrl(dbUrl);
			dataSource.setDriverClass("com.mysql.jdbc.Driver");
			dataSource.setInitialPoolSize(2);
			dataSource.setMinPoolSize(1);
			dataSource.setMaxPoolSize(10);
			dataSource.setMaxStatements(50);
			dataSource.setMaxIdleTime(60);
		} catch (PropertyVetoException e) {
			throw new RuntimeException(e);
		}
	}
	
	/**
	 * 该接口不支持事务。如果需要支持事务请使用getConn自行获取链接
	 * @return
	 */
	public static final QueryRunner createQueryRunner() {
		return new QueryRunner(dataSource);
	}
	
	/**
	 * 该接口支持事务。
	 * 自行conn.setAutocommit(false)，使用结束后conn.commit()
	 */
	public static final Connection getConn() {
		try {
			return dataSource.getConnection();
		} catch (SQLException ex) {
			ex.printStackTrace();
			return null;
		}
	}
}
