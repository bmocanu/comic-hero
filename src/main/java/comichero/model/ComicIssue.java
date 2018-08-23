package comichero.model;

public class ComicIssue {

    private String id;

    private String title;

    private String imageUrl;

    public ComicIssue(String id, String title, String imageUrl) {
        this.id = id;
        this.title = title;
        this.imageUrl = imageUrl;
    }

    public String getId() {
        return id;
    }

    public String getTitle() {
        return title;
    }

    public String getImageUrl() {
        return imageUrl;
    }
}
