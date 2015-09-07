import Ember from 'ember';

export default Ember.Component.extend({
    actions: {

        createTask: function(task) {
            this.sendAction('createTask', task);
            this.set('task', {});
        },

        deleteTask: function(task) {
            this.sendAction('deleteTask', task);
            this.set('task', {});
        }
    }
});
