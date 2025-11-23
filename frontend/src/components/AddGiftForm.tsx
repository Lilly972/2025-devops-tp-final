import { useForm } from '@tanstack/react-form';
import type { CreateGiftInput } from '../types';

interface AddGiftFormProps {
  personId: string;
  onSubmit: (data: CreateGiftInput) => void;
}

export function AddGiftForm({ personId, onSubmit }: AddGiftFormProps) {
  const form = useForm({
    defaultValues: {
      person_id: personId,
      title: '',
      description: '',
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
        🎁 Add a Gift Idea
      </h2>
      <div className="space-y-4">
        <form.Field name="title">
          {(field) => (
            <input
              type="text"
              className="christmas-input w-full"
              placeholder="Gift title..."
              value={field.state.value}
              onChange={(e) => field.handleChange(e.target.value)}
              required
            />
          )}
        </form.Field>
        <form.Field name="description">
          {(field) => (
            <textarea
              className="christmas-input w-full h-24 resize-none"
              placeholder="Gift description..."
              value={field.state.value}
              onChange={(e) => field.handleChange(e.target.value)}
              required
            />
          )}
        </form.Field>

        <button type="submit" className="christmas-button w-full">
          ➕ Add Gift Idea
        </button>
      </div>
    </form>
  );
}
