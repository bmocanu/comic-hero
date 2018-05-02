package comichero.retrievers;

import comichero.api.ComicIssue;

import java.time.LocalDate;

/**
 * Contract for a comic issue retriever. Upon command, it returns a new instance of a
 * {@link ComicIssue} for the given date.
 *
 * @author Bogdan Mocanu
 */
public interface Retriever {

    ComicIssue retrieveIssue(LocalDate date);

}
