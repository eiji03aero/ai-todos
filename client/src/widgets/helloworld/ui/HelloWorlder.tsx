import React, { useRef, useEffect, useState } from 'react'
import { Button } from '@/shared/ui/button'
import { Stopwatch } from '@/features/helloworld/lib/stopwatch'
import { FireworksAnimation } from '@/features/helloworld/lib/fireworks'

export const HelloWorlder: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const [stopwatch, setStopwatch] = useState<Stopwatch | null>(null)
  const [elapsedTime, setElapsedTime] = useState<string>('00:00:00')
  const fireworksRef = useRef<FireworksAnimation | null>(null)

  useEffect(() => {
    if (canvasRef.current) {
      fireworksRef.current = new FireworksAnimation(canvasRef.current)
    }
  }, [])

  const handleFireworks = () => {
    if (canvasRef.current && fireworksRef.current) {
      const rect = canvasRef.current.getBoundingClientRect()
      const x = rect.width / 2
      const y = rect.height / 2
      fireworksRef.current.trigger(x, y)
    }
  }

  const handleStartStopwatch = () => {
    if (!stopwatch) {
      const newStopwatch = new Stopwatch()
      newStopwatch.start((time: string) => {
        setElapsedTime(time)
      })
      setStopwatch(newStopwatch)
    } else {
      stopwatch.stop()
      setStopwatch(null)
    }
  }

  return (
    <div className="flex flex-col items-center justify-center space-y-4 p-4">
      <canvas
        ref={canvasRef}
        width={600}
        height={400}
        className="border border-gray-300 rounded-lg"
      />
      <div className="text-2xl font-bold mb-4">{elapsedTime}</div>
      <div className="flex space-x-4">
        <Button onClick={handleFireworks} variant="default">
          Great
        </Button>
        <Button 
          onClick={handleStartStopwatch} 
          variant={stopwatch ? 'destructive' : 'default'}
        >
          {stopwatch ? 'Stop' : 'Start stopwatch'}
        </Button>
      </div>
    </div>
  )
}