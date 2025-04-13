export class FireworksAnimation {
  private canvas: HTMLCanvasElement;
  private ctx: CanvasRenderingContext2D;
  private particles: Particle[] = [];

  constructor(canvas: HTMLCanvasElement) {
    this.canvas = canvas;
    this.ctx = canvas.getContext('2d')!;
  }

  private createParticles(x: number, y: number) {
    const particleCount = 50;
    const colors = ['#ff0000', '#00ff00', '#0000ff', '#ffff00', '#ff00ff'];

    for (let i = 0; i < particleCount; i++) {
      const angle = Math.random() * Math.PI * 2;
      const speed = Math.random() * 5 + 2;
      const color = colors[Math.floor(Math.random() * colors.length)];

      this.particles.push(new Particle(x, y, angle, speed, color));
    }
  }

  public trigger(x: number, y: number) {
    this.createParticles(x, y);
    this.animate();
  }

  private animate() {
    this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);

    for (let i = this.particles.length - 1; i >= 0; i--) {
      const particle = this.particles[i];
      particle.update();
      particle.draw(this.ctx);

      if (particle.alpha <= 0) {
        this.particles.splice(i, 1);
      }
    }

    if (this.particles.length > 0) {
      requestAnimationFrame(() => this.animate());
    }
  }
}

class Particle {
  x: number;
  y: number;
  private angle: number;
  private speed: number;
  private color: string;
  alpha: number;

  constructor(x: number, y: number, angle: number, speed: number, color: string) {
    this.x = x;
    this.y = y;
    this.angle = angle;
    this.speed = speed;
    this.color = color;
    this.alpha = 1;
  }

  update() {
    this.x += Math.cos(this.angle) * this.speed;
    this.y += Math.sin(this.angle) * this.speed + 0.2; // Gravity effect
    this.alpha -= 0.02;
    this.speed *= 0.95; // Slow down
  }

  draw(ctx: CanvasRenderingContext2D) {
    ctx.beginPath();
    ctx.arc(this.x, this.y, 2, 0, Math.PI * 2);
    ctx.fillStyle = this.color;
    ctx.globalAlpha = this.alpha;
    ctx.fill();
    ctx.closePath();
    ctx.globalAlpha = 1;
  }
}