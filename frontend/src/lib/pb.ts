import PocketBase, { RecordService } from 'pocketbase';
import type { SearchResponse } from './search';

export const baseUrl = import.meta.env.DEV
	? `http://${window.location.hostname}:8090`
	: window.location.origin;

export interface Base {
	id: string;
	created: string;
	updated: string;
}

interface TemplateSource {
	type: 'embed' | 'input' | 'custom';
	value: string;
}

interface TemplateResource {
	type: 'font' | 'image' | 'text';
	label: string;
	source: TemplateSource;
}

interface TemplateElement {
	type: 'text' | 'image';
	resource: string;
	position: {
		x: number;
		y: number;
	};
	size: {
		width: number;
		height: number;
	};
	options: {
		center: boolean;
		color: [number, number, number];
		font: {
			name: string;
			size: number;
		};
	};
}

export interface TemplateData {
	aspectRatio: [number, number];
	resources: Record<string, TemplateResource>;
	elements: TemplateElement[];
}

export interface Template extends Base {
	name: string;
	thumbnail: string;
	data: TemplateData;
}

export interface Image extends Base {
	name: string;
	tag: string;
	image: string;
}

interface TypedPocketBase extends PocketBase {
	collection(idOrName: string): RecordService;
	collection(idOrName: 'templates'): RecordService<Template>;
	collection(idOrName: 'images'): RecordService<Image>;
}

const pb = new PocketBase(baseUrl) as TypedPocketBase;

export default pb;

export interface RenderRequest {
	id: string;
	inputs: Record<string, unknown>;
	gap: number;
	margin: number;
	quantity: number;
	size: {
		width: number;
		height: number;
	};
	outline: boolean;
}

export const renderTemplate = async (request: RenderRequest) => {
	const res = await fetch(`${baseUrl}/render`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(request)
	});

	return res.blob();
};

export const searchImages = async (query: string, transparent = true, start = 1) => {
	const params = new URLSearchParams({
		q: query,
		start: start.toString(),
		transparent: transparent.toString()
	});
	const res = await pb.send(`/search?${params.toString()}`, {});

	return res as SearchResponse;
};
