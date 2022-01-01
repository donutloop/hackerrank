

struct box
{
	/**
	* Define three fields of type int: length, width and height
	*/

    int length;
    int width;
    int height;
};

typedef struct box box;

int get_volume(box b) {
	return b.height * b.length * b.width;
}

int is_lower_than_max_height(box b) {
	if (b.height < 41) {
        return 1;
    }
    return 0;
}

