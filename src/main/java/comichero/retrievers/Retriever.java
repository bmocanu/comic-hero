package comichero.retrievers;

import comichero.model.ComicIssue;

import java.time.LocalDate;

/**
 * Contract for a comic issue retriever. Upon command, it returns a new instance of a
 * {@link ComicIssue} for the given date.
 *
 * @author Bogdan Mocanu
 */
public interface Retriever {

    boolean isEnabled();

    ComicIssue retrieveIssue(LocalDate date);

}
