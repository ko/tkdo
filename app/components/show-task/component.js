import Ember from 'ember';

export default Ember.Component.extend({
    actions: {
        deleteTask: function(task) {
            this.sendAction('deleteTask', task);
            console.log('kenko:sent:' + task);
            this.set('task', {});
        }
    }
});
