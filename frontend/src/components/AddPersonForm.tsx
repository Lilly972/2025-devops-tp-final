import { useForm } from '@tanstack/react-form';
import type { CreatePersonInput } from '../types';

interface AddPersonFormProps {
  onSubmit: (data: CreatePersonInput) => void;
}

export function AddPersonForm({ onSubmit }: AddPersonFormProps) {
  const form = useForm({
    defaultValues: {
      name: '',
    },
    onSubmit: async ({ value }) => {
      onSubmit(value);
      form.reset();
    },
  });

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        e.stopPropagation();
        form.handleSubmit();
      }}
      className="christmas-card p-6"
    >
      <h2 className="text-2xl font-christmas font-bold mb-4 text-christmas-gold">
        🎅 Add a New Person
      </h2>
      <div className="space-y-4">
        <form.Field name="name">
          {(field) => (
            <input
              type="text"
              className="christmas-input w-full"
              placeholder="Enter person's name..."
              value={field.state.value}
              onChange={(e) => field.handleChange(e.target.value)}
              required
            />
          )}
        </form.Field>

        <button type="submit" className="christmas-button w-full">
          ➕ Add Person
        </button>
      </div>
    </form>
  );
}
