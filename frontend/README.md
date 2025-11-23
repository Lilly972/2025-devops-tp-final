# Frontend - Christmas Gift List

A modern React single-page application with TypeScript, TanStack Router, TanStack Query, and Tailwind CSS for managing Christmas gift lists.

## Tech Stack

- **React** 19
- **TypeScript** 5+
- **TanStack Router** - Type-safe routing
- **TanStack Query** - Data fetching and caching
- **TanStack Form** - Form management
- **Tailwind CSS** - Utility-first CSS framework
- **Vite** - Build tool and dev server
- **Vitest** - Unit testing framework
- **Storybook** - Component development environment

## Prerequisites

- Node.js 18 or higher
- npm (comes with Node.js)

## Getting Started

### 1. Install Dependencies

```bash
npm install
```

### 2. Start Development Server

```bash
npm run dev
```

The application will start at `http://localhost:5173`.

**Important**: The backend API must be running at `http://localhost:8080` for the application to work properly.

## API Proxy Configuration

### Development Mode

In development, Vite is configured to proxy all `/api` requests to the backend server. This is defined in `vite.config.ts`:

```typescript
server: {
  proxy: {
    '/api': {
      target: 'http://localhost:8080',
      changeOrigin: true,
    },
  },
}
```

This means:

- Frontend runs on `http://localhost:5173`
- Backend runs on `http://localhost:8080`
- All requests to `http://localhost:5173/api/*` are proxied to `http://localhost:8080/api/*`

### Production/Deployment Considerations

⚠️ **Important**: The Vite development proxy is **NOT** available in production builds!

When the frontend project builds it become a **static site** and loses all server habilities, including proxiing to the backend.

When deploying the frontend to production, **you MUST configure a redirect/proxy on your deployment platform** to route `/api` requests to your backend server. The frontend code cannot be modified, so this must be handled at the infrastructure level.

#### Required: Platform-Level Redirect/Proxy Configuration

Most modern hosting platforms provide built-in redirect or proxy capabilities. You need to configure your platform to proxy all `/api/*` requests to your backend API.

**Example Platforms and Their Redirect/Proxy Features:**

##### Netlify - `_redirects` file or `netlify.toml`

Create a `_redirects` file in your `public/` directory:

```
/api/*  https://your-backend-domain.com/api/:splat  200
```

Or use `netlify.toml`:

```toml
[[redirects]]
  from = "/api/*"
  to = "https://your-backend-domain.com/api/:splat"
  status = 200
  force = true
```

**Other Docker Reverse Proxy Options:**

- **[Nginx](https://hub.docker.com/_/nginx)** - High-performance web server and reverse proxy
- **[Traefik](https://hub.docker.com/_/traefik)** - Modern HTTP reverse proxy with automatic HTTPS
- **[Caddy](https://hub.docker.com/_/caddy)** - Easy-to-use web server with automatic HTTPS
- **[HAProxy](https://hub.docker.com/_/haproxy)** - Reliable, high-performance TCP/HTTP load balancer

## Available Commands

### Development

```bash
npm run dev              # Start Vite dev server (http://localhost:5173)
npm run build            # Build for production (outputs to dist/)
npm run preview          # Preview production build locally
```

### Testing

```bash
npm run test             # Run unit tests with Vitest
npm run test:ui          # Run tests with Vitest UI
npm run test:coverage    # Run tests with coverage report
```

### Code Quality

```bash
npm run lint             # Run ESLint
npm run lint:fix         # Fix ESLint errors automatically
npm run format           # Format code with Prettier
npm run format:check     # Check code formatting with Prettier
```

### Storybook

```bash
npm run storybook        # Start Storybook dev server (http://localhost:6006)
npm run build-storybook  # Build Storybook for production
```

## Project Structure

```
frontend/
├── public/                 # Static assets
├── src/
│   ├── assets/            # Images, fonts, etc.
│   ├── components/        # React components
│   │   ├── *.tsx          # Component implementation
│   │   ├── *.test.tsx     # Component tests
│   │   └── *.stories.tsx  # Storybook stories
│   ├── hooks/             # Custom React hooks
│   ├── lib/
│   │   └── api.ts         # API client functions
│   ├── pages/             # Page components
│   ├── routes/            # TanStack Router routes
│   │   ├── __root.tsx     # Root layout
│   │   ├── index.tsx      # Home page
│   │   └── person.$personId.tsx  # Person detail page
│   ├── test/
│   │   └── setup.ts       # Test configuration
│   ├── types/
│   │   └── index.ts       # TypeScript type definitions
│   ├── App.tsx            # Main app component
│   ├── App.css            # App styles
│   ├── main.tsx           # Application entry point
│   └── index.css          # Global styles
├── .storybook/            # Storybook configuration
├── vite.config.ts         # Vite configuration
├── vitest.config.ts       # Vitest configuration (if separate)
├── tsconfig.json          # TypeScript configuration
├── tailwind.config.js     # Tailwind CSS configuration
└── package.json           # Dependencies and scripts
```

## Building for Production

### Build the Application

```bash
npm run build
```

This creates an optimized production build in the `dist/` directory.

### Preview Production Build

```bash
npm run preview
```

This serves the production build locally for testing.

## Component Development

### Creating a New Component

1. Create the component file in `src/components/`
2. Create a test file `ComponentName.test.tsx`
3. Create a Storybook story `ComponentName.stories.tsx`

Example:

```typescript
// MyComponent.tsx
export function MyComponent({ title }: { title: string }) {
  return <div>{title}</div>;
}

// MyComponent.test.tsx
import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { MyComponent } from './MyComponent';

describe('MyComponent', () => {
  it('renders title', () => {
    render(<MyComponent title="Hello" />);
    expect(screen.getByText('Hello')).toBeInTheDocument();
  });
});

// MyComponent.stories.tsx
import type { Meta, StoryObj } from '@storybook/react';
import { MyComponent } from './MyComponent';

const meta: Meta<typeof MyComponent> = {
  component: MyComponent,
};

export default meta;
type Story = StoryObj<typeof MyComponent>;

export const Default: Story = {
  args: {
    title: 'Hello World',
  },
};
```

## Testing

### Unit Tests

Run unit tests with Vitest:

```bash
# Run all tests
npm run test

# Watch mode
npm run test -- --watch

# Coverage
npm run test:coverage
```

### Component Tests

Component tests use Vitest and React Testing Library. See existing tests in `src/components/*.test.tsx` for examples.

### Storybook Tests

Storybook stories can also be tested using the Vitest integration configured in this project.

## Styling

This project uses Tailwind CSS for styling. Key files:

- `tailwind.config.js` - Tailwind configuration
- `src/index.css` - Global styles and Tailwind imports
- Component files - Use Tailwind utility classes

## TypeScript

This is a fully typed TypeScript project. Type definitions are in:

- `src/types/index.ts` - Shared types
- Component files - Component-specific types

## Troubleshooting

### API Requests Fail in Development

- Ensure the backend is running at `http://localhost:8080`
- Check the proxy configuration in `vite.config.ts`
- Verify network requests in browser DevTools

### Production Build Not Working

- Test the build locally with `npm run preview`
- Ensure API routing is properly configured (reverse proxy or environment variables)
- Check browser console for errors
- Verify CORS settings if using cross-origin deployment

### Storybook Not Starting

- Clear node_modules and reinstall: `rm -rf node_modules && npm install`
- Check for port conflicts (default: 6006)

## Contributing

When contributing to the frontend:

1. Follow the existing code style
2. Write tests for new components
3. Create Storybook stories for new components
4. Run linting before committing: `npm run lint`
5. Ensure all tests pass: `npm run test`

## Additional Resources

- [React Documentation](https://react.dev/)
- [TanStack Router](https://tanstack.com/router)
- [TanStack Query](https://tanstack.com/query)
- [Tailwind CSS](https://tailwindcss.com/)
- [Vite](https://vitejs.dev/)
- [Vitest](https://vitest.dev/)
- [Storybook](https://storybook.js.org/)

## License

See root project README for license information
