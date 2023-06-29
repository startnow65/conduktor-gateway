'use client'

import { useState } from "react"
import { toast } from 'react-toastify';

const defaultObj = {
    name: '',
    email: '',
    favourite_food: '', // intentional spelling error!
    height: 0
}
export default function Home() {
    const [data, setDate] = useState(defaultObj as any);
    const [isLoading, setIsLoading] = useState(false);

    async function submitForm(e: { preventDefault: () => void }) {
        e.preventDefault()

        if (isLoading) return;

        if (!data.name?.length || !data.email?.length || !data.favourite_food?.length) {
            toast.error('Invalid input ...');
            return;
        }

        try {
            setIsLoading(true)
            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/publish`, {
                method: 'POST',
                body: JSON.stringify({
                    name: data.name,
                    email: data.email,
                    favourite_food: data.favourite_food,
                    height: +data.height
                }),
                headers: {
                    'Content-type': 'application/json; charset=UTF-8',
                },
            })
        } catch (e) {
            console.log(e);
        } finally {
            toast.success('Published!');
            setDate(defaultObj);
            setIsLoading(false);
        }
    }

    function handleChange(e: any) {
        e.preventDefault();
        setDate({ ...data, [e.target.name]: e.target.value });
    }

    return (
        <form onSubmit={submitForm}>
            <div className="space-y-12">
                <div className="border-b border-gray-900/10 pb-12">
                    <h2 className="text-base font-semibold leading-7 text-gray-900">User Information</h2>

                    <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
                        <div className="sm:col-span-3">
                            <label htmlFor="first-name" className="block text-sm font-medium leading-6 text-gray-900">
                                Full Name
                            </label>
                            <div className="mt-2">
                                <input
                                    type="text"
                                    name="name"
                                    id="name"
                                    value={data.name}
                                    onChange={handleChange}
                                    className="block w-full rounded-md border-0 px-2 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>

                        <div className="sm:col-span-3">
                            <label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900">
                                Email address
                            </label>
                            <div className="mt-2">
                                <input
                                    id="email"
                                    name="email"
                                    type="email"
                                    value={data.email}
                                    onChange={handleChange}
                                    autoComplete="email"
                                    className="block w-full rounded-md border-0 px-2 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>

                        <div className="sm:col-span-3">
                            <label htmlFor="favourite_food" className="block text-sm font-medium leading-6 text-gray-900">
                                Favorite Food
                            </label>
                            <div className="mt-2">
                                <input
                                    type="text"
                                    name="favourite_food"
                                    id="favourite_food"
                                    value={data.favourite_food}
                                    onChange={handleChange}
                                    autoComplete="favourite_food"
                                    className="block w-full rounded-md border-0 px-2 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>

                        <div className="sm:col-span-3">
                            <label htmlFor="height" className="block text-sm font-medium leading-6 text-gray-900">
                                Height
                            </label>
                            <div className="mt-2">
                                <input
                                    type="number"
                                    name="height"
                                    id="height"
                                    value={data.height}
                                    onChange={handleChange}
                                    autoComplete="height"
                                    className="block w-full rounded-md border-0 px-2 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div className="mt-6 flex items-center justify-end gap-x-6">
                <button
                    disabled={isLoading}
                    type="submit"
                    className="flex items-center justify-center rounded-md w-32 h-12 px-3 py-2 text-sm font-semibold text-white shadow-sm bg-indigo-600 hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                    Publish
                </button>
            </div>
        </form>
    )
}