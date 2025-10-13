type Transaction = {
  id: string;
  walletId: string;
  type: 'deposit' | 'withdrawal' | 'payment' | 'transfer';
  currency: string;
  description: string;
  createdAt: string;
};

export { type Transaction };
