package comichero.config;

/**
 * The model for each comic configuration fragment.
 *
 * @author Bogdan Mocanu
 */
public class ComicsItemNode {

    private boolean enabled = false;

    private boolean proxyImage = false;

    public boolean isEnabled() {
        return enabled;
    }

    public void setEnabled(boolean enabled) {
        this.enabled = enabled;
    }

    public boolean isProxyImage() {
        return proxyImage;
    }

    public void setProxyImage(boolean proxyImage) {
        this.proxyImage = proxyImage;
    }

    @Override
    public String toString() {
        return "ComicsNode{" +
                "enabled=" + enabled +
                ", proxyImage=" + proxyImage +
                '}';
    }
}
