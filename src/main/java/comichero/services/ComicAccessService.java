package comichero.services;

import comichero.api.ComicIssue;

import java.util.List;

/**
 * Provides access capabilities to comics and comic issues that are loaded in a store.
 *
 * @author Bogdan Mocanu
 */
public interface ComicAccessService {

    List<ComicIssue> getComicIssues(String... comicIds);

}
