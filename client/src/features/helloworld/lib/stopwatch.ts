export class Stopwatch {
  private startTime: number | null = null;
  private elapsedTime: number = 0;
  private isRunning: boolean = false;
  private updateCallback: ((time: string) => void) | null = null;
  private intervalId: number | null = null;

  start(callback?: (time: string) => void) {
    if (!this.isRunning) {
      this.startTime = Date.now();
      this.isRunning = true;
      
      if (callback) {
        this.updateCallback = callback;
        this.intervalId = setInterval(() => {
          if (this.updateCallback) {
            this.updateCallback(this.formatTime());
          }
        }, 1000) as unknown as number;
      }
    }
  }

  stop() {
    if (this.isRunning && this.startTime) {
      this.elapsedTime += Date.now() - this.startTime;
      this.startTime = null;
      this.isRunning = false;
      
      if (this.intervalId) {
        clearInterval(this.intervalId);
        this.intervalId = null;
      }
    }
  }

  reset() {
    this.startTime = null;
    this.elapsedTime = 0;
    this.isRunning = false;
    
    if (this.intervalId) {
      clearInterval(this.intervalId);
      this.intervalId = null;
    }
  }

  getElapsedTime(): { minutes: number; seconds: number } {
    let totalSeconds = this.elapsedTime;
    
    if (this.isRunning && this.startTime) {
      totalSeconds += Date.now() - this.startTime;
    }

    totalSeconds = Math.floor(totalSeconds / 1000);
    const minutes = Math.floor(totalSeconds / 60);
    const seconds = totalSeconds % 60;

    return { minutes, seconds };
  }

  formatTime(): string {
    const { minutes, seconds } = this.getElapsedTime();
    const hours = Math.floor(minutes / 60);
    const remainingMinutes = minutes % 60;
    
    return `${hours.toString().padStart(2, '0')}:${remainingMinutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
  }

  isActive(): boolean {
    return this.isRunning;
  }
}