import TheHeader from '@/components/TheHeader'
import './globals.css'
import { Inter } from 'next/font/google'

const inter = Inter({ subsets: ['latin'] })

export const metadata = {
    title: 'Kafka Gateway POC | Cloud Communication',
    description: 'Gateway POC',
}

export default function RootLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
            <body className={inter.className}>
                <TheHeader />
                <main className="max-w-5xl py-6 mx-auto px-4 sm:px-6 lg:px-8">
                    {children}
                </main>
            </body>
        </html>
    )
}
