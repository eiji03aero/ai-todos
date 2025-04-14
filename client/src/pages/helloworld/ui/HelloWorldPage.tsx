import React from 'react'
import { HelloWorlder } from '@/widgets/helloworld/ui/HelloWorlder'
import { TicTacToeBoard } from '@/features/tictactoe/ui/tictactoe-board'
import { useTicTacToe } from '@/features/tictactoe/model/use-tictactoe'

export const HelloWorldPage: React.FC = () => {
  const { gameState, makeMove, resetGame } = useTicTacToe();

  return (
    <div className="container mx-auto p-4">
      <div className="flex flex-col items-center justify-center space-y-6">
        <h1 className="text-4xl font-bold text-blue-600">Hello, World!</h1>
        <HelloWorlder />
        
        <div className="mt-8">
          <h2 className="text-3xl font-bold mb-6 text-center">Tic Tac Toe</h2>
          <TicTacToeBoard 
            gameState={gameState}
            onCellClick={makeMove}
            onReset={resetGame}
          />
        </div>
      </div>
    </div>
  )
}