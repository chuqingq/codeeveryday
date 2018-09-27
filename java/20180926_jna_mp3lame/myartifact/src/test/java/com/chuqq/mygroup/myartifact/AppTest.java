package com.chuqq.mygroup.myartifact;

import junit.framework.Test;
import junit.framework.TestCase;
import junit.framework.TestSuite;

import java.util.Date;
import java.util.concurrent.LinkedBlockingQueue;

/**
 * Unit test for simple App.
 */
public class AppTest 
    extends TestCase
{
    /**
     * Create the test case
     *
     * @param testName name of the test case
     */
    public AppTest( String testName )
    {
        super( testName );
    }

    /**
     * @return the suite of tests being tested
     */
    public static Test suite()
    {
        return new TestSuite( AppTest.class );
    }

    /**
     * Rigourous Test :-)
     */
    public void testApp()
    {
        assertTrue( true );
        System.out.println("123");
        Date now = new Date(); 
        System.out.println("now: " + now);
        System.out.println("" + System.currentTimeMillis());

        final LinkedBlockingQueue<Long> queue = new LinkedBlockingQueue<Long>();

        Thread child = new Thread() {
            public void run(){
                try {
                    Long i = queue.take();
                    System.out.println("child take: " + (System.nanoTime() - i));
                } catch (Exception e) {
                    //TODO: handle exception
                }
                
            }
        };
        child.start();

        try {
            Thread.sleep(3000);
        } catch (Exception e) {
            //TODO: handle exception
        }
        
        System.out.println("parent now: " + System.nanoTime());
        queue.offer((long)1024);

        try {
            child.join();
        } catch (Exception e) {
            //TODO: handle exception
        }
    }
}
