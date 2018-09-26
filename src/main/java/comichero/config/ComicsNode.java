package comichero.config;

/**
 * The "comics" node model from the application custom configuration. This class is a component
 * of the {@link Configuration}.
 *
 * @author Bogdan Mocanu
 */
public class ComicsNode {

    private ComicsItemNode sinfest = new ComicsItemNode();
    private ComicsItemNode dilbert = new ComicsItemNode();

    public ComicsItemNode getSinfest() {
        return sinfest;
    }

    public ComicsItemNode getDilbert() {
        return dilbert;
    }

    @Override
    public String toString() {
        return "ComicsNode{" +
                "sinfest=" + sinfest +
                ", dilbert=" + dilbert +
                '}';
    }
}
