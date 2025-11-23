import type { Person } from '../types';
import { Link } from '@tanstack/react-router';

interface PersonCardProps {
  person: Person;
  onDelete: (id: string) => void;
}

export function PersonCard({ person, onDelete }: PersonCardProps) {
  return (
    <div className="christmas-card p-6 hover:scale-105 transition-transform">
      <h3 className="text-2xl font-christmas font-bold mb-2 text-christmas-gold text-center">
        {person.name}
      </h3>
      <div className="flex gap-2 mt-4">
        <Link
          to="/person/$personId"
          params={{ personId: person.id }}
          className="christmas-button flex-1 text-center"
        >
          🎁 View Gifts
        </Link>
        <button
          onClick={() => onDelete(person.id)}
          className="bg-christmas-red/60 hover:bg-christmas-red text-white font-bold py-2 px-4 rounded-lg transition-all duration-200"
          aria-label={`Delete ${person.name}`}
        >
          🗑️
        </button>
      </div>
    </div>
  );
}
