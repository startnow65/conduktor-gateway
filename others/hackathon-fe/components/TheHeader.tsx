import Link from 'next/link'
// import Image from 'next/image'
// import logo from '@/assets/hellofresh-logo.svg'


export default function TheHeader() {
    return (
        <nav className="bg-gray-600">
            <div className="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex items-center justify-between h-16">
                    <div className="flex items-center">
                        <h1 className="text-gray-300 text-lg">HF Hackathon 2023</h1>
                    </div>
                    <div className="hidden md:block">
                        <div className="ml-10 flex items-baseline space-x-4">
                            <Link
                                href="/"
                                className="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                            >
                                Add Customer
                            </Link>
                            <Link
                                href="/delete"
                                className="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                            >
                                Forget User
                            </Link>
                            <Link
                                href="/list"
                                className="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                            >
                                List Customers
                            </Link>
                        </div>
                    </div>
                </div>
            </div>
        </nav>
    )
}