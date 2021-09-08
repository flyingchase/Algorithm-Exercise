package Exercises;

/**
 * @author :qlzhou;
 * @date: :创建于2021/9/612:34 下午
 */
public class broadcast {
    public static void main(String[] args) throws InterruptedException {
        for (int i = 0; i < 10; i++) {
            new Thread(() -> {
                try {
                    synchronized (broadcast.class) {
                        broadcast.class.wait();
                    }
                    System.out.println(Thread.currentThread().getName() + "done...");
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }

            }).start();
        }
        Thread.sleep(300);
        synchronized ( broadcast.class) {
            broadcast.class.notifyAll();
        }
    }
}
