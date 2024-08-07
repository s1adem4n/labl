export interface SearchResponse {
	kind: string;
	url: URL;
	queries: Queries;
	context: Context;
	searchInformation: SearchInformation;
	items: Item[];
}

interface URL {
	type: string;
	template: string;
}

interface Queries {
	request: Request[];
	nextPage: Request[];
}

interface Request {
	title: string;
	totalResults: string;
	searchTerms: string;
	count: number;
	startIndex: number;
	inputEncoding: string;
	outputEncoding: string;
	safe: string;
	cx: string;
}

interface Context {
	title: string;
}

interface SearchInformation {
	searchTime: number;
	formattedSearchTime: string;
	totalResults: string;
	formattedTotalResults: string;
}

export interface Item {
	kind: string;
	title: string;
	htmlTitle: string;
	link: string;
	displayLink: string;
	snippet: string;
	htmlSnippet: string;
	mime: string;
	fileFormat: string;
	image: Image;
}

interface Image {
	contextLink: string;
	height: number;
	width: number;
	byteSize: number;
	thumbnailLink: string;
	thumbnailHeight: number;
	thumbnailWidth: number;
}
