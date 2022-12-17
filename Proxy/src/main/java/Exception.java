import java.text.ParseException;

public class Exception extends ParseException {
    public Exception(String message, int where) {
        super(message, where);
    }
}
