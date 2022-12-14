package tamplon.snake.util.unsafe;

@FunctionalInterface
public interface UnsafeConsumer<T>
{

    void accept(final T value) throws Exception;
}
