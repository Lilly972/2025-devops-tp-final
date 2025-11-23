import type { Meta, StoryObj } from '@storybook/react';
import { PersonCard } from './PersonCard';

const meta: Meta<typeof PersonCard> = {
  title: 'Components/PersonCard',
  component: PersonCard,
  tags: ['autodocs'],
  argTypes: {
    onDelete: { action: 'deleted' },
  },
};

export default meta;
type Story = StoryObj<typeof PersonCard>;

export const Default: Story = {
  args: {
    person: {
      id: '1',
      name: 'Santa Claus',
      created_at: new Date().toISOString(),
    },
  },
};

export const LongName: Story = {
  args: {
    person: {
      id: '2',
      name: 'Ebenezer Scrooge the Third of Westminster',
      created_at: new Date().toISOString(),
    },
  },
};
