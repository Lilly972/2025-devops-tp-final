export interface Person {
  id: string;
  name: string;
  created_at: string;
}

export interface Gift {
  id: string;
  person_id: string;
  title: string;
  description: string;
  is_selected: boolean;
  created_at: string;
}

export interface CreatePersonInput {
  name: string;
}

export interface CreateGiftInput {
  person_id: string;
  title: string;
  description: string;
}

export interface UpdateGiftInput {
  title?: string;
  description?: string;
  is_selected?: boolean;
}
