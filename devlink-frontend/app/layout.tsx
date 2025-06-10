import '../styles/globals.css';
import { Inter } from 'next/font/google';
import { AuthProvider } from '@/context/auth';

const inter = Inter({ subsets: ['latin'] });

export const metadata = {
  title: 'DevLink',
};

import { ReactNode } from 'react';

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <AuthProvider>{children}</AuthProvider>
      </body>
    </html>
  );
}
