package comichero.api;

/**
 * A descriptor type for a comic. This includes a string ID, a name and a description.
 *
 * @author Bogdan Mocanu
 */
public class Comic {

    private String id;

    private String name;

    private String description;

    public Comic(String id, String name, String description) {
        this.id = id;
        this.name = name;
        this.description = description;
    }

    public String getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public String getDescription() {
        return description;
    }
}
