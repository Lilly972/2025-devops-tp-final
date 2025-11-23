import { createRootRoute, Outlet, Link } from '@tanstack/react-router';
import { TanStackRouterDevtools } from '@tanstack/router-devtools';
import { Snowflakes } from '../components/Snowflakes';

export const Route = createRootRoute({
  component: () => (
    <div className="min-h-screen">
      <Snowflakes />
      <header className="bg-christmas-red/80 backdrop-blur-sm border-b-4 border-christmas-gold shadow-lg">
        <div className="container mx-auto px-4 py-6">
          <div className="flex items-center justify-between">
            <Link to="/" className="flex items-center gap-3">
              <span className="text-5xl">🎄</span>
              <h1 className="text-4xl font-christmas font-bold text-christmas-gold">
                Christmas Gift List
              </h1>
            </Link>
            <nav className="flex gap-4">
              <Link
                to="/"
                className="christmas-button"
                activeProps={{ className: 'ring-2 ring-christmas-gold' }}
              >
                🏠 Home
              </Link>
            </nav>
          </div>
        </div>
      </header>
      <main className="container mx-auto px-4 py-8">
        <Outlet />
      </main>
      <TanStackRouterDevtools />
    </div>
  ),
});
