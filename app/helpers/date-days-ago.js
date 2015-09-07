import Ember from 'ember';

export function dateDaysAgo(params/*, hash*/) {
    let date = new Date(params.get(0));
    let today = new Date();
    let dateAgoMs = Math.abs(today - date);
    let dateAgoDays = Math.round(dateAgoMs / 3600 / 24 / 1000);

    if (dateAgoDays === 0) {
        return "Today";
    } else if (dateAgoDays === 1) {
        return "1 day ago";
    } else {
        return dateAgoDays + " days ago";
    }
}

export default Ember.Helper.helper(dateDaysAgo);
