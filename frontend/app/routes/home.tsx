import { Link } from "react-router";
import type { Route } from "./+types/home";

export function meta({ }: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

type Author = {
  ID: string;
  Name: string;
  CreatedAt: number;
};

export async function clientLoader({
  params,
}: Route.ClientLoaderArgs) {
  let data = await fetch(`/api/authors`);
  return data.json() as Promise<Author[]>;
}

export default function Home({ loaderData }: Route.ComponentProps) {
  console.log("loaderData", loaderData);
  return <div style={{ padding: '20px' }}>
    <ul>
      {loaderData.map((author) => (
        <li key={author.ID}>{author.Name}</li>
      ))}
    </ul>
    <Link to="/about">About</Link>
  </div>;
}
