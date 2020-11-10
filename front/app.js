// Copyright (c) 2020, Arveto Ink. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

const stdsubjets = {
	'tomas': {
		id: 'tomas',
		pseudo: 'Tomas Benedito Bonito',
		email: 'tomas@fai.mil',
		level: 'visitor',
		bot: false,
	},
	'zech': {
		id: 'zech',
		pseudo: 'Zech',
		email: 'zech@fai.mil',
		level: 'standard',
		bot: false,
	},
	'tzu': {
		id: 'tzu',
		pseudo: 'Tzu Shikaï',
		email: 'tzu@fai.mil',
		level: 'administrator',
		bot: false,
	},
	'r2d2': {
		id: 'bot-r2d2',
		pseudo: 'R2D2',
		email: 'r2d2@bot',
		level: 'standard',
		bot: true,
	},
	'custom': {}
};

document.addEventListener('DOMContentLoaded', () => {
	document.querySelectorAll('#list button').forEach(b => {
		const s = stdsubjets[b.id];
		b.addEventListener('click', () => {
			s.teams = getTeams();
			s.app = getApp();
			generate(s);
		});
	});

	const search = new URLSearchParams(document.location.search);
	const app = search.get('app');
	if (app) document.getElementById('app').value = app;
	const r = search.get('r');
	if (r) document.getElementById('r').value = r;

	let updateAppID = [];
	let updateAppURL = [];
	let updateRedirect = [];
	let updateTeams = [];

	const custom = document.getElementById('custom');
	document.getElementById('customGen').addEventListener('click',
		() => generate(stdsubjets.custom));
	['id', 'pseudo', 'email', 'level', 'bot'].map(id => document.getElementById(id))
		.forEach(input => {
			stdsubjets.custom[input.id] = input.value;
			input.addEventListener('input', () => {
				stdsubjets.custom[input.id] = input.value;
				custom.href = '/redirect?' + new URLSearchParams(stdsubjets.custom);
			});
		});

	[custom, ...document.querySelectorAll('#list a')]
	.forEach(a => {
		a.addEventListener('click', event => {
			event.preventDefault();
			window.open(a.href).focus();
		});

		let s = stdsubjets[a.id];

		function update() {
			a.href = '/redirect?' + new URLSearchParams(s);
		}

		function updater(ar, f) {
			ar.push(v => {
				f(v);
				update();
			})
		}

		updater(updateAppID, id => s.app = id);
		updater(updateAppURL, u => s.u = u);
		updater(updateRedirect, r => s.r = r);
		updater(updateTeams, t => s.teams = t);
	});

	// Teams
	function udpatingTeams() {
		const teams = getTeams();
		updateTeams.forEach(u => u(teams));
	}
	udpatingTeams();
	document.getElementById('teams').addEventListener('input', udpatingTeams);

	function updator(array, id) {
		const input = document.getElementById(id);

		function up() {
			const v = input.value;
			array.forEach(u => u(v));
		}
		input.addEventListener('input', up);
		up();
	}
	updator(updateAppID, 'app');
	updator(updateAppURL, 'u');
	updator(updateRedirect, 'r');

	document.querySelectorAll('input').forEach(input => {
		input.addEventListener('input', hideJWTOut);
	});
}, {
	once: true,
});

function getTeams() {
	return document.getElementById('teams').value
		.split(',')
		.map(t => t.trim())
		.filter(t => t.length)
		.sort();
}

function getApp() {
	return document.getElementById('app').value;
}

function redirectURL(query) {
	return '/redirect' + new URLSearchParams(query);
}

// Fetch the generated JWT.
async function generate(query) {
	const out = document.getElementById('out');
	out.value = 'loading ...';
	out.hidden = false;

	const jwt = await fetch('/generate?' + new URLSearchParams(query))
		.then(rep => rep.text())
		.then(text => text.trim());
	out.value = jwt;

	out.onclick = () => {
		const i = document.createElement('input');
		document.body.appendChild(i);
		i.type = 'text';
		i.value = jwt;
		i.select();
		document.execCommand('copy');
		i.remove();

		out.classList.add('copied');
		setTimeout(() => out.classList.remove('copied'), 500);
	};
}

function hideJWTOut() {
	document.getElementById('out').hidden = true;
}
