import type { Meta, StoryObj } from '@storybook/react';
import { GiftCard } from './GiftCard';

const meta: Meta<typeof GiftCard> = {
  title: 'Components/GiftCard',
  component: GiftCard,
  tags: ['autodocs'],
  argTypes: {
    onDelete: { action: 'deleted' },
    onSelect: { action: 'selected' },
  },
};

export default meta;
type Story = StoryObj<typeof GiftCard>;

export const Default: Story = {
  args: {
    gift: {
      id: '1',
      person_id: '1',
      title: 'New Sleigh',
      description: 'A modern, eco-friendly sleigh with GPS navigation',
      is_selected: false,
      created_at: new Date().toISOString(),
    },
  },
};

export const Selected: Story = {
  args: {
    gift: {
      id: '2',
      person_id: '1',
      title: 'Comfortable Boots',
      description: 'Warm, waterproof boots for chimney climbing',
      is_selected: true,
      created_at: new Date().toISOString(),
    },
  },
};

export const LongDescription: Story = {
  args: {
    gift: {
      id: '3',
      person_id: '1',
      title: 'Premium Cookie Recipe Collection',
      description:
        'An extensive collection of cookie recipes from around the world, including traditional favorites and modern innovations. Features step-by-step instructions with beautiful photography.',
      is_selected: false,
      created_at: new Date().toISOString(),
    },
  },
};
