import type { User, Wallet, Card, Transaction } from '../shared/types';

/**
 * /src/database/dummy_data.ts
 *
 * Two dummy users, each with one wallet. Wallets include cards and transactions.
 * Uses shared types from shared/types.
 */

export type DummyUser = User & { password: string };
export const dummy_users: DummyUser[] = [
  {
    id: 'user_1_9f1a2b3c',
    name: 'Alice Santos',
    email: 'alice@example.com',
    password: 'password123',
    createdAt: '2024-01-15T09:30:00.000Z',
  },
  {
    id: 'user_2_a7c4d5e6',
    name: 'Bob Oliveira',
    email: 'bob@example.com',
    password: 'securepass',
    createdAt: '2024-02-02T14:20:00.000Z',
  },
];

export const dummy_wallets: Wallet[] = [
  {
    id: 'wallet_1_111aaa',
    userId: 'user_1_9f1a2b3c',
    currency: 'USD',
    balance: 1265.0,
    createdAt: '2024-01-15T09:30:00.000Z',
    cards: [
      {
        id: 'card_1_visa_4242',
        brand: 'visa',
        last4: '4242',
        expiryDate: '12/26',
        cardNumber: '4242424242424242',
        entity: 'Bank of Examples',
        cvc: '123',
        cardHolder: 'Alice Santos',
        createdAt: '2024-01-15T09:35:00.000Z',
      } as Card,
    ],
    transactions: [
      {
        id: 'tx_1_deposit_1000',
        walletId: 'wallet_1_111aaa',
        type: 'deposit',
        currency: 'USD',
        description: 'Initial deposit',
        createdAt: '2024-01-15T09:31:00.000Z',
      } as Transaction,
      {
        id: 'tx_2_deposit_300',
        walletId: 'wallet_1_111aaa',
        type: 'deposit',
        currency: 'USD',
        description: 'Paycheck',
        createdAt: '2024-01-30T08:00:00.000Z',
      } as Transaction,
      {
        id: 'tx_3_payment_grocery',
        walletId: 'wallet_1_111aaa',
        cardId: 'card_1_visa_4242',
        type: 'payment',
        currency: 'USD',
        description: 'Grocery store',
        merchant: 'Green Market',
        createdAt: '2024-02-02T12:45:00.000Z',
      } as Transaction,
    ],
  },
  {
    id: 'wallet_2_222bbb',
    userId: 'user_2_a7c4d5e6',
    currency: 'EUR',
    balance: 480.75,
    createdAt: '2024-02-02T14:20:00.000Z',
    cards: [
      {
        id: 'card_2_master_5555',
        brand: 'mastercard',
        last4: '5555',
        expiryDate: '07/25',
        cardNumber: '5555555555555555',
        entity: 'Example Credit Union',
        cvc: '456',
        cardHolder: 'Bob Oliveira',
        createdAt: '2024-02-02T14:25:00.000Z',
      } as Card,
    ],
    transactions: [
      {
        id: 'tx_4_deposit_500',
        walletId: 'wallet_2_222bbb',
        type: 'deposit',
        currency: 'EUR',
        description: 'Transfer from savings',
        createdAt: '2024-02-02T14:21:00.000Z',
      } as Transaction,
      {
        id: 'tx_5_payment_cafe',
        walletId: 'wallet_2_222bbb',
        cardId: 'card_2_master_5555',
        type: 'payment',

        currency: 'EUR',
        description: 'Coffee and snack',
        merchant: 'Café Bairro',
        createdAt: '2024-02-03T09:10:00.000Z',
      } as Transaction,
      {
        id: 'tx_6_withdrawal_atm',
        walletId: 'wallet_2_222bbb',
        type: 'withdrawal',

        currency: 'EUR',
        description: 'ATM withdrawal',
        createdAt: '2024-02-04T11:00:00.000Z',
      } as Transaction,
    ],
  },
];
