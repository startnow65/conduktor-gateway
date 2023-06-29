'use client'

import { useEffect, useState } from "react";

export default async function Home() {
    const [users, setUsers] = useState([])
    const [loading, setLoading] = useState(false)

    useEffect(() => {
        setLoading(true)
        fetch(`${process.env.NEXT_PUBLIC_API_URL}/consume`)
            .then(response => response.json())
            .then(json => setUsers(json))
            .finally(() => {
                setLoading(false)
            })
    }, [])

    return (
        <ul role="list" className="divide-y divide-gray-100">
            {Boolean(users.length) && users.map((person: any) => (
                <li key={`${person.email}-${Math.random()}`} className="flex justify-between gap-x-6 py-5">
                    <div className="flex gap-x-4">
                        <div className="min-w-0 flex-auto">
                            <p className="text-sm font-semibold leading-6 text-gray-900">Name: {person.name}</p>
                            <p className="mt-1 truncate text-xs leading-5 text-gray-500">Email: {person.email}</p>
                        </div>
                    </div>
                    <div className="hidden sm:flex sm:flex-col w-[200px]">
                        <p className="text-sm leading-6 text-gray-900">Favorite Food: {person.favourite_food}</p>
                        <p className="mt-1 text-xs leading-5 text-gray-500">
                            Height: <time dateTime={person.lastSeenDateTime}>{person.height}</time>
                        </p>
                    </div>
                </li>
            ))}
        </ul>
    );
}