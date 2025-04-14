export type Player = 'X' | 'O';

export type BoardCell = Player | null;

export type Board = [
  [BoardCell, BoardCell, BoardCell],
  [BoardCell, BoardCell, BoardCell],
  [BoardCell, BoardCell, BoardCell]
];

export type GameResult = 
  | { status: 'ongoing' }
  | { status: 'draw' }
  | { status: 'win', winner: Player };

export interface TicTacToeGameState {
  board: Board;
  currentPlayer: Player;
  result: GameResult;
}