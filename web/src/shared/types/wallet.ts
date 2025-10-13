import type { Card } from './card';
import { type Transaction } from './transaction';

type Wallet = {
  id: string;
  userId: string;
  balance: number;
  currency: string;
  createdAt: string;
  cards: Card[];
  transactions: Transaction[];
};

export { type Wallet };
