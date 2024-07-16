import pb, { type Image, type Template } from './pb';

export const isAuthenticated = () => authenticated;

let authenticated = $state(pb.authStore.isValid);
pb.authStore.onChange(() => {
	authenticated = pb.authStore.isValid;
});

export const templates: Template[] = $state([]);
export const images: Image[] = $state([]);

export const load = async () => {
	const templatesData = await pb.collection('templates').getList(0, 999);
	templates.splice(0, templates.length, ...templatesData.items);

	const imagesData = await pb.collection('images').getList(0, 999);
	images.splice(0, images.length, ...imagesData.items);
};

export const subscribe = async () => {
	const templatesUnsubscribe = await pb.collection('templates').subscribe('*', (e) => {
		if (e.action === 'create') {
			templates.push(e.record);
		} else if (e.action === 'update') {
			const index = templates.findIndex((t) => t.id === e.record.id);
			if (index !== -1) {
				templates[index] = e.record;
			}
		} else if (e.action === 'delete') {
			const index = templates.findIndex((t) => t.id === e.record.id);
			if (index !== -1) {
				templates.splice(index, 1);
			}
		}
	});

	const imagesUnsubscribe = await pb.collection('images').subscribe('*', (e) => {
		if (e.action === 'create') {
			images.push(e.record);
		} else if (e.action === 'update') {
			const index = images.findIndex((t) => t.id === e.record.id);
			if (index !== -1) {
				images[index] = e.record;
			}
		} else if (e.action === 'delete') {
			const index = images.findIndex((t) => t.id === e.record.id);
			if (index !== -1) {
				images.splice(index, 1);
			}
		}
	});

	return () => {
		templatesUnsubscribe();
		imagesUnsubscribe();
	};
};
