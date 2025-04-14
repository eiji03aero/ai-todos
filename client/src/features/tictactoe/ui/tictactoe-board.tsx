import React from 'react';
import { TicTacToeGameState } from '@/entities/tictactoe/model/game-state';

interface TicTacToeBoardProps {
  gameState: TicTacToeGameState;
  onCellClick: (row: number, col: number) => void;
  onReset: () => void;
}

export function TicTacToeBoard({ 
  gameState, 
  onCellClick, 
  onReset 
}: TicTacToeBoardProps) {
  const renderCell = (row: number, col: number) => {
    const cellValue = gameState.board[row][col];
    return (
      <div 
        key={`${row}-${col}`}
        className={`
          w-24 h-24 border border-gray-300 flex items-center justify-center 
          text-6xl font-bold cursor-pointer transition-colors
          ${cellValue ? 'text-blue-600' : 'text-gray-400'}
          hover:bg-gray-100
        `}
        onClick={() => onCellClick(row, col)}
      >
        {cellValue || ''}
      </div>
    );
  };

  const renderGameStatus = () => {
    const { result, currentPlayer } = gameState;
    
    if (result.status === 'win') {
      return `Winner: ${result.winner}`;
    }
    
    if (result.status === 'draw') {
      return 'Draw!';
    }

    return `Current Player: ${currentPlayer}`;
  };

  return (
    <div className="flex flex-col items-center justify-center space-y-4">
      <div className="grid grid-cols-3 gap-2">
        {gameState.board.map((row, rowIndex) => 
          row.map((_, colIndex) => renderCell(rowIndex, colIndex))
        )}
      </div>
      <div className="text-2xl font-semibold">
        {renderGameStatus()}
      </div>
      {(gameState.result.status !== 'ongoing') && (
        <button 
          onClick={onReset}
          className="
            px-4 py-2 bg-blue-500 text-white rounded 
            hover:bg-blue-600 transition-colors
          "
        >
          Reset Game
        </button>
      )}
    </div>
  );
}