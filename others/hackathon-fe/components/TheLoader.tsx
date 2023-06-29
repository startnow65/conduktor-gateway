import Image from 'next/image';
import LoaderImage from '../assets/loader.svg';

export default function Loader() {
    return (
        <Image alt='loader' width={45} src={LoaderImage} />
    )
}