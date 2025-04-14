import { useState, useCallback } from 'react';
import { TicTacToeGameState, Player, Board, BoardCell, GameResult } from '@/entities/tictactoe/model/game-state';

export function useTicTacToe() {
  const initialBoard: Board = [
    [null, null, null],
    [null, null, null],
    [null, null, null]
  ];

  const [gameState, setGameState] = useState<TicTacToeGameState>({
    board: initialBoard,
    currentPlayer: 'X',
    result: { status: 'ongoing' }
  });

  const makeMove = useCallback((row: number, col: number) => {
    if (
      gameState.result.status !== 'ongoing' || 
      gameState.board[row][col] !== null
    ) return;

    const newBoard = gameState.board.map((r, rowIndex) => 
      r.map((cell, colIndex) => 
        rowIndex === row && colIndex === col ? gameState.currentPlayer : cell
      )
    ) as Board;

    const newResult = checkGameResult(newBoard);
    const nextPlayer: Player = gameState.currentPlayer === 'X' ? 'O' : 'X';

    setGameState({
      board: newBoard,
      currentPlayer: nextPlayer,
      result: newResult
    });
  }, [gameState]);

  const resetGame = useCallback(() => {
    setGameState({
      board: initialBoard,
      currentPlayer: 'X',
      result: { status: 'ongoing' }
    });
  }, []);

  return {
    gameState,
    makeMove,
    resetGame
  };
}

function checkGameResult(board: Board): GameResult {
  // 勝利判定
  const winningCombos = [
    // 横
    [[0,0], [0,1], [0,2]],
    [[1,0], [1,1], [1,2]],
    [[2,0], [2,1], [2,2]],
    // 縦
    [[0,0], [1,0], [2,0]],
    [[0,1], [1,1], [2,1]],
    [[0,2], [1,2], [2,2]],
    // 斜め
    [[0,0], [1,1], [2,2]],
    [[0,2], [1,1], [2,0]]
  ];

  for (const combo of winningCombos) {
    const [first, second, third] = combo;
    const [a, b, c] = combo.map(([r, c]) => board[r][c]);
    
    if (a && a === b && a === c) {
      return { status: 'win', winner: a };
    }
  }

  // 引き分け判定
  const isDraw = board.every(row => row.every(cell => cell !== null));
  if (isDraw) {
    return { status: 'draw' };
  }

  return { status: 'ongoing' };
}