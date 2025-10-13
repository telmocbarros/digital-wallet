import { motion } from 'framer-motion';

export default function BluewaveBankCard() {
  return (
    <motion.div
      className="w-80 p-4 rounded-2xl shadow-lg bg-gradient-to-r from-blue-600 to-blue-400 text-white"
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
    >
      <div className="flex justify-between items-center mb-4">
        <div className="text-xl font-bold">Bluewave Bank</div>
        <div className="text-sm">**** 1234</div>
      </div>
      <div className="text-3xl font-semibold mb-2">$5,000</div>
      <div className="flex justify-center mt-4">
        <div className="w-12 h-12 bg-white rounded-lg flex items-center justify-center text-black font-mono">
          QR
        </div>
      </div>
    </motion.div>
  );
}
