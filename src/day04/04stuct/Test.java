class Point {
    private int x;
    private int y;

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public int getX() {
        return x;
    }
}

public class Test {
    public double dist(Point a, Point b) {
        return Math.sqrt(Math.pow(a.getX() - b.getX(), 2) + Math.pow(a.getY() - b.getY(), 2));
    }

    public static void main(String[] args) {
        Point a = new Point(2, 2);
        Point b = new Point(3, 3);
        double result = new Test().dist(a, b);
        System.out.println(result);
    }
}