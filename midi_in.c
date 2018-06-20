#include <stdio.h>
#include <alsa/asoundlib.h>
#include <alsa/rawmidi.h>
#include <signal.h>

int stop=0;
void sighandler(int dum) {
	stop = 1;
}

int main() {
	int err;
	snd_rawmidi_t *handle_in = 0;
	
	err = snd_rawmidi_open(&handle_in, NULL, "hw:1,0,0", 0);
	if (err) {
		fprintf(stderr,"snd_rawmidi_open failed: %d\n",err);
	}
		
	signal(SIGINT, sighandler);
	
	if (handle_in) {
		fprintf(stderr,"Read midi in\n");
		fprintf(stderr,"Press ctrl-c to stop\n");
	}

	if (handle_in) {
		unsigned char ch;
		while (!stop) {
			snd_rawmidi_read(handle_in,&ch,1);
			fprintf(stderr,"read %02x\n",ch);
		}
	}
	
	if (handle_in) {
		snd_rawmidi_drain(handle_in); 
		snd_rawmidi_close(handle_in);   
	}

	return 0;
}
