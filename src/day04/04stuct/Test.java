class Point {
    public int x;
    public int y;

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public String toString() {
        return "x: " + x + ",y: " + y;
    }
}

public class Test {
    // java中拷贝的是对象的引用
    public void dist(Point a) {
        a.x = 2;
        a.y = 2;
    }

    public static void main(String[] args) {
        Point A = new Point(0, 0);
        System.out.println(A);
        new Test().dist(A);
        System.out.println(A);
        // output:
        // x: 0,y: 0
        // x: 2,y: 2

    }
}