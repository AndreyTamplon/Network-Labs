package tamplon.snake.gui.menu;

import tamplon.snake.gamelogic.gameobjects.config.Config;
import tamplon.snake.gui.game.SnakesGameView;

import java.net.InetSocketAddress;
import java.util.function.Consumer;

@FunctionalInterface
public interface GameJoiner
{

    void joinGame(
            final String playerName,
            final Config gameConfig,
            final InetSocketAddress hostAddress,
            final SnakesGameView gameView,
            final Runnable onSuccess,
            final Consumer<String> onError);
}
