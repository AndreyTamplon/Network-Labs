package tamplon.snake.gui.menu;

import tamplon.snake.gamelogic.gameobjects.config.Config;
import tamplon.snake.gui.game.SnakesGameView;

@FunctionalInterface
public interface GameStarter
{

    void startGame(
            final String playerName,
            final Config gameConfig,
            final SnakesGameView gameView);
}
